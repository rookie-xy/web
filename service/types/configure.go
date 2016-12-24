
/*
 * Copyright (C) 2016 Meng Shi
 */


package types


const (
    ConfNoargs = 0x00000001
    ConfTake1 = 0x00000002
    ConfTake2 = 0x00000004
    ConfTake3 = 0x00000008
    ConfTake4 = 0x00000010
    ConfTake5 = 0x00000020
    ConfTake6 = 0x00000040
    ConfTake7 = 0x00000080
    ConfTake12 = (ConfTake1|ConfTake2)

    ConfMore1 = 0x00000800
    ConfMore2 = 0x00001000
    ConfAny = 0x00000400

    ConfBlock = 0x00000100
    ConfFlag = 0x00000200

    MainConf = 0x01000000
    AnyConf = 0x0F000000
    DirectConf = 0x00010000
)


type Configure struct {
    File    string

    Action  interface {
        Set()
        Get()
        Parse()
    }
}
