require_relative './node'
require_relative './binary_tree_errors'

class BinaryTree
  include BinaryTreeErrors

  def initialize(root)
    @root = root
    @size = 0
  end

  def insert_left(node, data)
    # O(1)
    raise NodeOverride if node.left
    node.left = Node.new(node, data)
    @size += 1
    node.left.data
  end

  def insert_right(node, data)
    # O(1)
    raise NodeOverride if node.right
    node.right = Node.new(node, data)
    @size += 1
    node.right.data
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
  end

  def self.merge(left, right, data=nil)
    # O(1)
    raise NilMerge unless left && right
    new_root = Node.new(nil, data)
    new_root.left = left.root
    new_root.right = right.root
    merged_tree = BinaryTree.new(new_root)
    merged_tree.size = left.size + right.size
    merged_tree
  end
end