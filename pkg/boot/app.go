package boot

import (
	starter2 "pronghorn/pkg/starter"
)

func init() {

	starter2.Register(&starter2.PropsStarter{})
	starter2.Register(&starter2.LBStarter{})
	starter2.Register(&starter2.ProxyStarter{})

}
