require 'rspec'
require_relative '../binary_tree'
require_relative '../node'
require_relative '../binary_tree_errors'

RSpec.describe BinaryTree do

  def create_btree
    root_node = Node.new(nil, 'test data')
    BinaryTree.new(root_node)
  end

  left_test_string = 'left'
  right_test_string = 'right'

  context 'initialize' do
    it 'initializes with root node having data equal to root_data argument' do
      test_string = 'test data'
      btree = BinaryTree.new(test_string)
      expect(btree.root.data).to eq(test_string)
    end

    it 'initializes with size 1' do
      btree = create_btree
      expect(btree.size).to eq(1)
    end
  end

  context 'insert_left' do
    it 'inserts data to the left of a node' do
      btree = create_btree
      btree.insert_left(btree.root, left_test_string)
      expect(btree.root.left.data).to eq(left_test_string)
    end

    it 'raises NodeOverride error if the target node already has a left child' do
      btree = create_btree
      btree.insert_left(btree.root, left_test_string)
      expect{ btree.insert_left(btree.root, left_test_string) }.to raise_error(BinaryTreeErrors::NodeOverride)
    end

    it 'returns the node containing the inserted data' do
      btree = create_btree
      node = btree.insert_left(btree.root, left_test_string)
      expect(node).to eq(btree.root.left)
    end
  end

  context 'insert_right' do
    it 'inserts data to the right of a node' do
      btree = create_btree
      btree.insert_right(btree.root, right_test_string)
      expect(btree.root.right.data).to eq(right_test_string)
    end

    it 'raises NodeOverride error if the target node already has a right child' do
      btree = create_btree
      btree.insert_right(btree.root, right_test_string)
      expect{ btree.insert_right(btree.root, right_test_string) }.to raise_error(BinaryTreeErrors::NodeOverride)
    end

    it 'returns the node containing the inserted data' do
      btree = create_btree
      node = btree.insert_right(btree.root, right_test_string)
      expect(node).to eq(btree.root.right)
    end
  end

  context 'remove_left' do
    it 'removes the left branch from the given node' do
      btree = create_btree
      first_left_child = btree.insert_left(btree.root, left_test_string)
      first_right_child = btree.insert_right(btree.root, right_test_string)
      btree.insert_left(first_left_child, left_test_string)
      btree.insert_right(first_left_child, right_test_string)
      btree.insert_left(first_left_child.left, left_test_string)
      btree.remove_left(btree.root)
      expect(btree.size).to eq(2)
      expect(btree.root.left).to be(nil)
    end

    it 'returns the root of the removed branch' do
      btree = create_btree
      first_left_child = btree.insert_left(btree.root, left_test_string)
      btree.insert_left(first_left_child, left_test_string)
      expect(btree.remove_left(first_left_child)).to eq(first_left_child)
    end
  end

  context 'remove_right' do
    it 'removes the right branch from the given node' do
      btree = create_btree
      first_right_child = btree.insert_right(btree.root, right_test_string)
      first_left_child = btree.insert_left(btree.root, left_test_string)
      btree.insert_right(first_right_child, right_test_string)
      btree.insert_left(first_right_child, left_test_string)
      btree.insert_right(first_right_child.right, right_test_string)
      btree.remove_right(btree.root)
      expect(btree.size).to eq(2)
      expect(btree.root.right).to eq(nil)
    end

    it 'returns the root of the removed branch' do
      btree = create_btree
      first_right_child = btree.insert_right(btree.root, right_test_string)
      btree.insert_right(first_right_child, right_test_string)
      expect(btree.remove_right(first_right_child)).to eq(first_right_child)
    end
  end

  context 'merge' do
    it 'merges two binary trees into a new binary tree' do
      left = create_btree
      right = create_btree
      merged_btree = BinaryTree.merge(left, right)
      expect(merged_btree.root).to have_attributes(left: left.root, right: right.root, data: nil)
    end

    it 'raises ArgumentError if left is not a binary tree' do
      right = create_btree
      expect{ BinaryTree.merge(nil, right) }.to raise_error(ArgumentError)
    end

    it 'raises ArgumentError if right is not a binary tree' do
      left = create_btree
      expect { BinaryTree.merge(left, nil) }.to raise_error(ArgumentError)
    end
  end

end