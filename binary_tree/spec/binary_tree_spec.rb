require 'rspec'
require_relative '../binary_tree'
require_relative '../node'
require_relative '../binary_tree_errors'

RSpec.describe BinaryTree do

  def create_btree
    root_node = Node.new(nil, 'test data')
    BinaryTree.new(root_node)
  end

  context 'initialize' do
    it 'initializes with root attribute set to root argument' do
      root_node = Node.new(nil, 'test data')
      btree = BinaryTree.new(root_node)
      expect(btree.root).to eq(root_node)
    end

    it 'initializes with size 0' do
      btree = create_btree
      expect(btree.size).to eq(0)
    end
  end

  context 'insert_left' do
    test_string = 'left'

    it 'inserts data to the left of a node' do
      btree = create_btree
      btree.insert_left(btree.root, test_string)
      expect(btree.root.left.data).to eq(test_string)
    end

    it 'raises NodeOverride error if the target node already has a left child' do
      btree = create_btree
      btree.insert_left(btree.root, test_string)
      expect{ btree.insert_left(btree.root, test_string) }.to raise_error(BinaryTreeErrors::NodeOverride)
    end

    it 'returns the node containing the inserted data' do
      btree = create_btree
      node = btree.insert_left(btree.root, test_string)
      expect(node).to eq(btree.root.left)
    end
  end

  context 'insert_right' do
    test_string = 'right'

    it 'inserts data to the right of a node' do
      btree = create_btree
      btree.insert_right(btree.root, test_string)
      expect(btree.root.right.data).to eq(test_string)
    end

    it 'raises NodeOverride error if the target node already has a right child' do
      btree = create_btree
      btree.insert_right(btree.root, test_string)
      expect{ btree.insert_right(btree.root, test_string) }.to raise_error(BinaryTreeErrors::NodeOverride)
    end

    it 'returns the node containing the inserted data' do
      btree = create_btree
      node = btree.insert_right(btree.root, test_string)
      expect(node).to eq(btree.root.right)
    end
  end
end