package ctci

import (
	"fmt"
	"testing"
)

func TestFindRoute(t *testing.T) {
	a := NewDirectedGraphNode("a")
	b := NewDirectedGraphNode("b")
	c := NewDirectedGraphNode("c")
	d := NewDirectedGraphNode("d")
	e := NewDirectedGraphNode("e")
	f := NewDirectedGraphNode("f")
	g := NewDirectedGraphNode("g")

	nilTestCases := []struct {
		s, t *DirectedGraphNode
	}{
		{s: nil, t: nil},
		{s: a, t: nil},
		{s: nil, t: b},
	}

	for _, tc := range nilTestCases {
		t.Run(fmt.Sprintf("FindRouteBirdirectional(%v, %v) returns nil", tc.s, tc.t), func(t *testing.T) {
			got := FindRoute(tc.s, tc.t)
			if got != nil {
				t.Errorf("want nil, got %+v", got)
			}
		})
	}

	t.Run("returns a slice with s as the only element when s == t", func(t *testing.T) {
		got := FindRoute(a, a)
		if len(got) > 1 || got[0] != a {
			fmt.Printf("got %+v", got)
		}
	})

	t.Run("finds route through linear graph", func(t *testing.T) {
		a.children = append(a.children, b)
		b.children = append(b.children, c)

		want := "a -> b -> c"
		got := RouteToString(FindRoute(a, c))
		if got != want {
			t.Errorf("want route %s, got %s", want, got)
		}
	})

	t.Run("finds shortest route", func(t *testing.T) {
		c.children = append(c.children, []*DirectedGraphNode{d, e}...)
		d.children = append(d.children, e)
		b.children = append(b.children, f)
		g.children = append(g.children, d)

		want := "a -> b -> c -> e"
		got := RouteToString(FindRoute(a, e))
		if got != want {
			t.Errorf("want route %s, got %s", want, got)
		}
	})

	t.Run("returns nil when no route exists", func(t *testing.T) {
		f.children = nil
		got := FindRoute(a, g)
		if got != nil {
			t.Errorf("want nil slice, got %+v", got)
		}
	})

	t.Run("terminates correctly when graph is cyclical", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic in cyclical graph: %v", r)
			}
		}()

		b.children = make([]*DirectedGraphNode, 1)
		b.children[0] = c
		c.children = make([]*DirectedGraphNode, 1)
		c.children[0] = a
		FindRoute(a, g)
	})
}

func TestFindRouteBidirectional(t *testing.T) {
	a := NewBirdirecGraphNode("a")
	b := NewBirdirecGraphNode("b")
	c := NewBirdirecGraphNode("c")
	d := NewBirdirecGraphNode("d")
	e := NewBirdirecGraphNode("e")
	f := NewBirdirecGraphNode("f")
	g := NewBirdirecGraphNode("g")

	nilTestCases := []struct {
		s, t *BidirecGraphNode
	}{
		{s: nil, t: nil},
		{s: a, t: nil},
		{s: nil, t: b},
	}

	for _, tc := range nilTestCases {
		t.Run(fmt.Sprintf("FindRouteBirdirectional(%v, %v) returns nil", tc.s, tc.t), func(t *testing.T) {
			got := FindRouteBidirectional(tc.s, tc.t)
			if got != nil {
				t.Errorf("want nil, got %+v", got)
			}
		})
	}

	t.Run("returns a slice with s as the only element when s == t", func(t *testing.T) {
		got := FindRouteBidirectional(a, a)
		if len(got) > 1 || got[0] != a {
			fmt.Printf("got %+v", got)
		}
	})

	t.Run("finds route through linear graph", func(t *testing.T) {
		a.children = append(a.children, b)
		b.children = append(b.children, c)
		b.parents = append(b.parents, a)
		c.parents = append(c.parents, b)

		want := "a -> b -> c"
		got := BidirecRouteToString(FindRouteBidirectional(a, c))
		if got != want {
			t.Errorf("want route %s, got %s", want, got)
		}
	})

	t.Run("finds shortest route", func(t *testing.T) {
		c.children = append(c.children, []*BidirecGraphNode{d, e}...)
		d.children = append(d.children, e)
		d.parents = append(d.parents, c)
		e.parents = append(e.parents, []*BidirecGraphNode{d, c}...)
		b.children = append(b.children, f)
		f.parents = append(f.parents, b)
		g.children = append(g.children, d)
		g.parents = append(g.parents, f)

		want := "a -> b -> c -> e"
		got := BidirecRouteToString(FindRouteBidirectional(a, e))
		if got != want {
			t.Errorf("want route %s, got %s", want, got)
		}
	})

	t.Run("returns nil when no route exists", func(t *testing.T) {
		f.children = nil
		g.parents = nil
		got := FindRouteBidirectional(a, g)
		if got != nil {
			t.Errorf("want nil slice, got %+v", got)
		}
	})

	t.Run("terminates correctly when graph is cyclical", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("panic in cyclical graph: %v", r)
			}
		}()

		b.children = make([]*BidirecGraphNode, 1)
		b.children[0] = c
		c.children = make([]*BidirecGraphNode, 1)
		c.children[0] = a
		a.parents = append(a.parents, c)
		FindRouteBidirectional(a, g)
	})
}
