package direction

// Alias hide the real type of the enum 
// and users can use it to define the var for accepting enum 
type Alias = int

type list struct { 
    N Alias
    E Alias
    S Alias
    W Alias
}

// Enum for public use
var Enum = &list{ 
    N: 0,
    E: 1,
    S: 2,
    W: 3,
}