require_relative './node'

class AVLTree
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
      if node.left.right.height > node.left.left.height
        return rotate_left_right(node)
      else
        return rotate_right(node)
      end
    elsif right_heavy?(node)
      if node.right.left.height > node.left.left.height
        return rotate_right_left(node)
      else
        return rotate_left(node)
      end
    end

    node
  end

  private def left_heavy?(node)
    lh = node.left ? node.left.height : 0
    rh = node.right ? node.right.height : 0
    lh - rh == 2
  end

  private def right_heavy?(node)
    lh = node.left ? node.left.height : 0
    rh = node.right ? node.right.height : 0
    rh - lh == 2
  end

  private def rotate_right(original_parent_node)
    new_parent = original_parent_node.left
    original_parent_node.left = new_parent.right
    new_parent.right = original_parent_node

    set_height(original_parent_node)
    set_height(new_parent)

    new_parent
  end

  private def rotate_left(original_parent_node)
    new_parent = original_parent_node.right
    original_parent_node.right = new_parent.left
    new_parent.left = original_parent_node

    set_height(original_parent_node)
    set_height(new_parent)

    new_parent
  end

  private def rotate_left_right(original_parent_node)
    original_parent_node.left = rotate_left(original_parent_node.left)
    return rotate_right(original_parent_node)
  end

  private def rotate_right_left(original_parent_node)
    original_parent_node.right = rotate_right(original_parent_node.right)
    return rotate_left(original_parent_node)
  end

  private def set_height(node)
    # O(1)
    lh = node.left ? node.left.height : 0
    rh = node.right ? node.right.height : 0
    max = lh > rh ? lh : rh
    node.height = max + 1
  end

  private def remove(key)
    search(key).deleted = true
  end
end