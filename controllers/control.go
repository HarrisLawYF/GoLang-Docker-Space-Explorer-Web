package controllers

import (
	"strconv"
	"github.com/astaxie/beego"
	"SpaceApp/controllers/direction"
	"SpaceApp/controllers/robot"
	"SpaceApp/controllers/ground"
)

type edge = ground.Edge
type machine = robot.Robot
type RobotController struct {
	beego.Controller
}

func (c *RobotController) Get() {
	c.Data["display"] = true
	c.Data["robotx"] = "N/A"
	c.Data["roboty"] = "N/A"
	c.Data["mapg"] = "N/A"
	c.Data["cdirection"] = "N/A"
	c.Data["ccoordinate"] = "N/A"
	c.TplName = "result.html"
}

func (c *RobotController) Post() {
	map_size, err := c.GetInt("map_size")
	pos1, err := c.GetInt("x_pos")
	pos2, err := c.GetInt("y_pos")
	
	if(err != nil){
		c.Data["error"] = err.Error
		c.Data["display"] = false
	}
	
	if(pos1 < 0 || pos1 > map_size){
		pos1 = 0
	}
	
	if(pos2 < 0 || pos2 > map_size){
		pos2 = 0
	}
	
	ground := ground.Create(map_size)
	navigator := robot.Create(pos1,pos2,direction.Enum.E)
	
	//Set the values for use in the template
	c.Data["display"] = true
	c.Data["robotx"] = navigator.Get_x()
	c.Data["roboty"] = navigator.Get_y()
	c.Data["mapg"] = ground
	c.Data["cdirection"] = navigator.Get_direction_str()
	c.Data["ccoordinate"] = strconv.Itoa(navigator.Get_x()) + ","+ strconv.Itoa(navigator.Get_y())
	inst := make(chan string)
	Find_shortest_path(ground, navigator.Get_x(), navigator.Get_y(), map_size, &navigator, inst)
	c.Data["instructions"] = inst
    c.TplName = "result.html"
}

func Find_shortest_path(edges [][]edge, current_x int, current_y int, map_bound_size int, bot *machine, inst chan string){
	go func() {
		if(current_x == map_bound_size && current_y == map_bound_size){
			inst <- "Destination reached: (X: " + strconv.Itoa(current_x) + ", Y: " + strconv.Itoa(current_y)+")"
			close(inst)
		} else if(current_x == map_bound_size){
			next_x := current_x
			next_y := current_y+1
			reply_bot := Move_robot(bot, next_x, next_y)
			inst <- reply_bot + ": "+"Robot moving to (X: "+strconv.Itoa(next_x)+", Y: "+strconv.Itoa(next_y)+")"
			Find_shortest_path(edges, next_x, next_y, map_bound_size, bot, inst)
		} else if(current_y == map_bound_size){
			next_x := current_x+1
			next_y := current_y
			reply_bot := Move_robot(bot, next_x, next_y)
			inst <- reply_bot + ": "+"Robot moving to (X: "+strconv.Itoa(next_x)+", Y: "+strconv.Itoa(next_y)+")"
			Find_shortest_path(edges, next_x, next_y, map_bound_size, bot, inst)
		} else{
			next_x, next_y := Get_min(edges[current_x+1][current_y], edges[current_x][current_y+1], current_x+1, current_y, current_x, current_y+1)
			reply_bot := Move_robot(bot, next_x, next_y)
			inst <- reply_bot + ": "+"Robot moving to (X: "+strconv.Itoa(next_x)+", Y: "+strconv.Itoa(next_y)+")"
			Find_shortest_path(edges, next_x, next_y, map_bound_size, bot, inst)
		}
	}()

	
}

func Get_min(edge1 edge, edge2 edge, edge1_x int, edge1_y int, edge2_x int, edge2_y int)(int,int){
	x:=0
	y:=0
	if(edge1.Weight < edge2.Weight){
		x = edge1_x
		y = edge1_y
	} else{
		x = edge2_x
		y = edge2_y
	}
	return x,y
}

func Move_robot(bot *machine, next_x int, next_y int)(string){
	diff_x := next_x - bot.Get_x()
	diff_y := next_y - bot.Get_y()
	switch{
		case diff_x == 1 && diff_y == 0:
			if(bot.Get_direction() == 0){
				*bot = bot.Turn_right()
				*bot = bot.Move_forward()
				return "Turn right, move forward"
			} else if(bot.Get_direction() == 1 || bot.Get_direction() == -3){
				*bot = bot.Move_forward()
				return "Move forward"
			} else if(bot.Get_direction() == 2 || bot.Get_direction() == -2){
				*bot = bot.Turn_left()
				*bot = bot.Move_forward()
				return "Turn left, move forward"
			} else{
				*bot = bot.Turn_right()
				*bot = bot.Turn_right()
				*bot = bot.Move_forward()
				return "Turn right, turn right, move forward"
			}
		case diff_x == 0 && diff_y == 1:
			if(bot.Get_direction() == 0){
				*bot = bot.Move_forward()
				return "Move forward"
			} else if(bot.Get_direction() == 1 || bot.Get_direction() == -3){
				*bot = bot.Turn_left()
				*bot = bot.Move_forward()
				return "Turn left, move forward"
			} else if(bot.Get_direction() == 2 || bot.Get_direction() == -2){
				*bot = bot.Turn_left()
				*bot = bot.Turn_left()
				*bot = bot.Move_forward()
				return "Turn left, turn left, move forward"
			} else{
				*bot = bot.Turn_right()
				*bot = bot.Move_forward()
				return "Turn right, move forward"
			}
		case diff_x == 0 && diff_y == 0:
			return "Robot moving to (x,y): " + strconv.Itoa(next_x)+", "+strconv.Itoa(next_y)
		default:
			return "Unknown"
	}
}
