package robot

type Robot struct {  
    x 			int
    y 			int
    direction	int
}

func Create(x int, y int, direction int)(Robot){
	bot := Robot{x,y,direction}
	return bot
}

func (bot Robot) Turn_left()(Robot) {  
	bot.direction = (bot.direction - 1) % 4
	return bot
}

func (bot Robot) Turn_right()(Robot) {  
	bot.direction = (bot.direction + 1) % 4
	return bot
}

func (bot Robot) Get_x()(int) {  
	return bot.x
}

func (bot Robot) Get_y()(int) {  
	return bot.y
}

func (bot Robot) Get_direction()(int) {  
	return bot.direction
}

func (bot Robot) Get_direction_str()(string) { 
	if(bot.direction == 0){
		return "North"
	} else if(bot.direction == 1 || bot.direction == -3){
		return "East"
	} else if(bot.direction == 2 || bot.direction == -2){
		return "South"
	} else{
		return "West"
	}
}

func (bot Robot) Move_forward()(Robot) {  
    if(bot.direction == 0){
		bot.y+=1
	} else if(bot.direction == 1 || bot.direction == -3){
		bot.x+=1
	} else if(bot.direction == 2 || bot.direction == -2){
		bot.y-=1
	} else if(bot.direction == 3 || bot.direction == -1){
		bot.x-=1
	}
	return bot
}