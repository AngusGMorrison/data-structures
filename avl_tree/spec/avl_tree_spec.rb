require 'rspec'
require_relative '../avl_tree';

RSpec.describe AVLTree do

  context 'initialize' do
    it 'initializes an AVL tree with root nil' do
      tree = AVLTree.new
      expect(tree.root).to be(nil);
    end
  end

  context 'insert' do
    left_heavy_tree = AVLTree.new

    it 'makes the first node inserted the root node' do
      left_heavy_tree.insert(6, "six")
      expect(left_heavy_tree.root.key).to be(6);
    end

    it 'inserts lower nodes to the left of the root node' do
      left_heavy_tree.insert(3, "three")
      expect(left_heavy_tree.root.left.key).to be(3)
    end

    it 'inserts higher nodes to the right of the root node' do
      left_heavy_tree.insert(9, "nine")
      expect(left_heavy_tree.root.right.key).to be(9)
    end

    it 'right-rotates the tree when the left branch grows too long with a long left tail' do
      left_heavy_tree.insert(2, "two")
      left_heavy_tree.insert(1, "one")
      expect(left_heavy_tree.root.left.key).to be(2)
      expect(left_heavy_tree.root.left.right.key).to be(3)
      expect(left_heavy_tree.root.left.left.key).to be(1)
    end

    it 'left-right rotates the tree when the left branch grows too long with a long right tail' do
      left_heavy_tree.insert(4, "four")
      expect(left_heavy_tree.root.key).to be(3)
      expect(left_heavy_tree.root.right.key).to be(6)
      expect(left_heavy_tree.root.right.left.key).to be(4)
      expect(left_heavy_tree.root.right.right.key).to be(9)
      expect(left_heavy_tree.root.left.key).to be(2)
      expect(left_heavy_tree.root.left.left.key).to be(1)
    end

    right_heavy_tree = AVLTree.new
    right_heavy_tree.insert(6, "six")
    right_heavy_tree.insert(3, "three")
    right_heavy_tree.insert(9, "nine")

    it 'left-rotates the tree when the right branch grows too long with a long right tail' do
      right_heavy_tree.insert(11, "eleven")
      right_heavy_tree.insert(13, "thirteen")
      expect(right_heavy_tree.root.right.key).to be(11)
      expect(right_heavy_tree.root.right.right.key).to be(13)
      expect(right_heavy_tree.root.right.left.key).to be(9)
    end

    it 'right-left rotates the tree when the right branch grows too long with a long left tail' do
      right_heavy_tree.insert(8, "eight")
      expect(right_heavy_tree.root.key).to be(9)
      expect(right_heavy_tree.root.right.key).to be(11)
      expect(right_heavy_tree.root.right.right.key).to be(13)
      expect(right_heavy_tree.root.left.key).to be(6)
      expect(right_heavy_tree.root.left.left.key).to be(3)
      expect(right_heavy_tree.root.left.right.key).to be(8)
    end
  end

  context 'search' do
    tree = AVLTree.new
    tree.insert(6, "six")
    tree.insert(3, "three")
    tree.insert(9, "nine")


    it 'returns the node with the given key' do
      node = tree.search(9)
      expect(node.data).to eq("nine");
    end

    it 'returns nil if the node is not found' do
      node = tree.search(11)
      expect(node).to be(nil)
    end

    it 'returns nil if the node is deleted' do
      tree.remove(3)
      node = tree.search(3)
      expect(node).to be(nil)
    end
  end

  context 'remove' do
    tree = AVLTree.new
    tree.insert(6, "six")
    tree.insert(3, "three")
    tree.insert(9, "nine")

    it 'takes a key to remove and marks the corresponding node as deleted' do
      node = tree.remove(3)
      expect(node.deleted).to be(true)
    end

    it 'returns nil if the key isn\'t found in the tree' do
      expect(tree.remove(20)).to be(nil)
    end
  end

end