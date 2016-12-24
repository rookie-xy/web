
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


var (
    Ok    =  0
    Error = -1
    Again = -2
)


const (
    CORE_MODULE = 0x45524F43
    SYSTEM_MODULE = 0x464E4F43
    USER_MODULE = 0x464E4F43
)


type InitModuleFunc func(cycle *Cycle) uint
type InitRoutineFunc func(cycle *Cycle) uint


type Module struct {
    CtxIndex      uint
    Index         uint
    Context      *Context
    Commands      []Command
    Type          uint
    InitModule    InitModuleFunc
    InitRoutine   InitRoutineFunc
}


func (i *Module) CoreInit(modules []*Module) (*Option, error) {
    var cycle *Cycle
    var m  int

    for m = 0; modules[m] != nil; m++ {
        mod := modules[m]

        if mod.Type != CORE_MODULE {
            continue
        }

        mod.InitModule(cycle)

        mod.InitRoutine(cycle)
    }

    return nil, nil
}


func (i *Module) SystemInit(option *Option) *Cycle {
    return nil
}


func (i *Module) UserInit(cycle *Cycle) {
}
