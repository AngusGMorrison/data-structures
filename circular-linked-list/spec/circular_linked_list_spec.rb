require 'rspec'
require_relative '../circular_linked_list'

RSpec.describe CircularLinkedList do

  def n_item_list(n)
    list = CircularLinkedList.new
    for i in 1..n do
      list.insert(i)
    end
    list
  end

  context 'initialization' do
    it 'is initialized with length 0' do
      list = CircularLinkedList.new
      expect(list.length).to eq(0)
    end

    it 'is initialized with head nil' do
      list = CircularLinkedList.new
      expect(list.head).to be(nil)
    end

    it 'is initialized with current nil' do
      list = CircularLinkedList.new
      expect(list.current).to be(nil)
    end
  end

  context 'insert' do
    it 'allows a new node to become the head and head.next values of an empty list' do
      list = n_item_list(1)
      expect(list.head).to have_attributes(data: 1, next: list.head)
    end

    it 'allows new nodes to be added to the "end" of the circle' do
      list = n_item_list(2)
      expect(list.last_node).to have_attributes(data: 2, next: list.head)
    end

    it 'increases length correctly as the list grows' do
      list = n_item_list(2)
      expect(list.length).to eq(2)
    end
  end

  context 'delete' do
    it 'deletes body nodes' do
      list = n_item_list(2)
      list.delete(list.head.next)
      expect(list.length).to eq(1)
    end

    it 'deletes head node' do
      list = n_item_list(2)
      original_tail = list.last_node
      list.delete(list.head)
      expect(list.head).to eq(original_tail)
    end

    it 'correctly connects previous node to next node on deletion' do
      list = CircularLinkedList.new
      node1 = list.insert(1)
      node2 = list.insert(2)
      node3 = list.insert(3)
      list.delete(node2)
      expect(node1.next).to eq(node3)
    end
  end

  context 'concat' do
    it 'concatenates two linked lists' do
      list = n_item_list(2)
      list.concat(n_item_list(3))
      expect(list.length).to eq(5)
    end

    it 'joins the second list between the last node and the head of the receiver' do
      list1 = n_item_list(2)
      original_list1_tail = list1.last_node
      list2 = n_item_list(2)
      original_list2_head = list2.head
      original_list2_tail = list2.last_node

      list1.concat(list2)
      expect(original_list1_tail.next).to eq(original_list2_head)
      expect(original_list2_tail.next).to eq(list1.head)
    end

    it 'raises an ArgumentError if the argument is not a CircularLinkedList' do
      list = n_item_list(2)
      expect { list.concat(1) }.to raise_error(ArgumentError)
    end
  end

  context 'clear' do
    it 'deletes all nodes from the list' do
      list = n_item_list(10)
      list.clear
      expect(list.length).to eq(0)
    end
  end

  context 'find' do
    it 'returns a node for which the block evaluates to true' do
      list = n_item_list(2)
      expect(list.find { |node| node.data == 2 }).to eq(list.last_node)
    end

    it 'returns the first node for which the block evaluates to true' do
      list = CircularLinkedList.new
      2.times { list.insert(1) }
      expect(list.find { |node| node.data == 1}).to eq(list.head)
    end

    it 'returns nil if a matching node is not found' do
      list = n_item_list(5)
      expect(list.find { |node| node.data == 6 }).to be(nil)
    end

    it 'raises a LocalJumpError if no block is given' do
      list = n_item_list(2)
      expect { list.find }.to raise_error(LocalJumpError)
    end
  end

  context 'each' do
    it 'iterates over each node in the list' do
      list = n_item_list(5)
      tracker = []
      list.each { |node| tracker << node.data }
      expect(tracker).to eq([1, 2, 3, 4, 5])
    end

    it 'raises a LocalJumpError if no block is given' do
      list = n_item_list(1)
      expect { list.each }.to raise_error(LocalJumpError)
    end
  end

  context 'to_a' do
    it 'returns the list\'s values in an array' do
      list = n_item_list(3)
      expect(list.to_a).to eq([1, 2, 3])
    end
  end

  context 'map' do
    it 'returns a copy of the list if no block is given' do
      list1 = n_item_list(3)
      list2 = list1.map
      expect(list2.to_a).to eq(list1.to_a)
      expect(list2).not_to eq(list1)
    end

    it 'returns a new list derived from the receiver and the given block' do
      list1 = n_item_list(3)
      list2 = list1.map { |value| value * 2 }
      expect(list2.to_a).to eq([2, 4, 6])
    end
  end


end