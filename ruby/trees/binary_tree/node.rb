class Node
  attr_accessor :parent, :data, :left, :right

  def initialize(parent_node, data)
    @parent = parent_node,
    @data = data
    @left = nil
    @right = nil
  end
end