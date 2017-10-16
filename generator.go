package generator

// YieldFunc is the yield function
type YieldFunc func(interface{})

// GenerateFunc is the generator function
type GenerateFunc func(YieldFunc)

// Generator is the generator type
type Generator struct {
	channel chan interface{}
	running bool
}

// MakeGenerator constructs a Generator instance
func MakeGenerator(generate GenerateFunc) (g *Generator) {
	g = &Generator{
		channel: make(chan interface{}),
		running: true,
	}
	go func() {
		defer g.Stop()
		generate(func(val interface{}) {
			if g.running {
				g.channel <- val
				<-g.channel
			}
		})
	}()
	return
}

// Next gets the next generated value
func (g *Generator) Next() (next interface{}, ok bool) {
	if next, ok = <-g.channel; ok {
		g.channel <- true
	}
	return
}

// Stop stops the generator
func (g *Generator) Stop() {
	if g.running {
		g.running = false
		close(g.channel)
	}
}
