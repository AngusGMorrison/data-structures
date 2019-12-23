require_relative './node'
require_relative './binary_tree_errors'

class BinaryTree
  include BinaryTreeErrors

  attr_reader :root
  attr_accessor :size

  def initialize(root_data=nil)
    @root = Node.new(nil, root_data)
    @size = 1
  end

  def insert_left(node, data)
    # O(1)
    raise NodeOverride if node.left
    node.left = Node.new(node, data)
    @size += 1
    node.left
  end

  def insert_right(node, data)
    # O(1)
    raise NodeOverride if node.right
    node.right = Node.new(node, data)
    @size += 1
    node.right
  end

  def remove_left(node)
    # O(n)
    if node.left
      remove_left(node.left)
      remove_right(node.left)
      node.left = nil
      @size -= 1
    end
    node
  end

  def remove_right(node)
    # O(n)
    if node.right
      remove_right(node.right)
      remove_left(node.right)
      node.right = nil
      @size -= 1
    end
    node
  end

  def self.merge(left, right, data=nil)
    # O(1)
    unless left.is_a?(BinaryTree) && right.is_a?(BinaryTree)
      raise ArgumentError.new("Expected two binary tress")
    end
    merged_tree = BinaryTree.new(data)
    merged_tree.root.left = left.root
    merged_tree.root.right = right.root
    merged_tree.size = left.size + right.size
    merged_tree
  end
end