require_relative './node'

class AVLTree
  attr_reader :root

  def initialize
    @root = nil
  end

  def insert(key, data=nil)
    # O(log2n)
    @root = insert_and_balance(@root, key, data)
  end

  private def insert_and_balance(node, key, data)
    return Node.new(key, data) unless node

    if key < node.key
      node.left = insert_and_balance(node.left, key, data)
    elsif key > node.key
      node.right = insert_and_balance(node.right, key, data)
    else
      node.data = data
      node.deleted = false
    end

    balance(node)
  end

  private def balance(node)
    set_height(node)
    if left_heavy?(node)
      return long_right_branch?(node) ? rotate_left_right(node) : rotate_right(node)
    elsif right_heavy?(node)
      return long_left_branch?(node) ? rotate_right_left(node) : rotate_left(node)
    end
    node
  end

  private def left_heavy?(node)
    height(node.left) - height(node.right) == 2
  end

  private def right_heavy?(node)
    height(node.right) - height(node.left) == 2
  end

  private def long_right_branch?(node)
    return false unless node.left
    height(node.left.right) > height(node.left.left)
  end

  private def long_left_branch?(node)
    return false unless node.right
    height(node.right.left) > height(node.right.right)
  end

  private def rotate_left_right(original_parent_node)
    original_parent_node.left = rotate_left(original_parent_node.left)
    return rotate_right(original_parent_node)
  end

  private def rotate_right_left(original_parent_node)
    original_parent_node.right = rotate_right(original_parent_node.right)
    return rotate_left(original_parent_node)
  end

  private def rotate_left(original_parent_node)
    new_parent = original_parent_node.right
    original_parent_node.right = new_parent.left
    new_parent.left = original_parent_node

    set_height(original_parent_node)
    set_height(new_parent)

    new_parent
  end

  private def rotate_right(original_parent_node)
    new_parent = original_parent_node.left
    original_parent_node.left = new_parent.right
    new_parent.right = original_parent_node

    set_height(original_parent_node)
    set_height(new_parent)

    new_parent
  end

  private def set_height(node)
    # O(1)
    lh = height(node.left)
    rh = height(node.right)
    max = lh > rh ? lh : rh
    node.height = max + 1
  end

  private def height(node)
    return 0 unless node
    node.height
  end

  def remove(key)
    # O(log2n)
    node = search(key)
    node.deleted = true if node
    node
  end

  def search(key)
    # O(log2n)
    node = search_from_node(@root, key)
    node unless node.deleted
  end

  def search_from_node(node, key)
    # O(log2n)
    return nil unless node
    if key < node.key
      return search_from_node(node.left, key)
    elsif key > node.key
      return search_from_node(node.right, key)
    end
    node
  end
end