package framework

type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)
	Put(string, ...ControllerHandler)
	Delete(string, ...ControllerHandler)

	Group(string) IGroup

	Use(middlewares ...ControllerHandler)
}

type Group struct {
	core   *Core
	prefix string

	parent *Group

	middlewares []ControllerHandler
}

func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:        core,
		prefix:      prefix,
		parent:      nil,
		middlewares: []ControllerHandler{},
	}
}

func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(g.getAbsolutePrefix()+uri, allHandlers...)
}

func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(g.getAbsolutePrefix()+uri, allHandlers...)
}

func (g *Group) Put(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Put(g.getAbsolutePrefix()+uri, allHandlers...)
}

func (g *Group) Delete(uri string, handlers ...ControllerHandler) {
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Delete(g.getAbsolutePrefix()+uri, allHandlers...)
}

func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = middlewares
}

func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}

	return append(g.parent.getMiddlewares(), g.middlewares...)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}

	return g.parent.getAbsolutePrefix() + g.prefix
}

func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}
