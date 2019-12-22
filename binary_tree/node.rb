class Node
  attr_accessor :parent, :data

  def initialize(parent_node, data)
    @parent = parent_node,
    @data = data
end