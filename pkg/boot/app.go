package boot

import "pronghorn/pkg/starter"

func init() {

	starter.Register(&starter.PropsStarter{})
	starter.Register(&starter.LBStarter{})
	starter.Register(&starter.ProxyStarter{})

}
