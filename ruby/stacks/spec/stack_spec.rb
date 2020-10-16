require 'rspec'
require_relative '../stack'

RSpec.describe Stack do

  def n_item_stack(n)
    stack = Stack.new
    n.times { |i| stack.push(i + 1) }
    stack
  end

  context 'initialize' do
    it 'initializes with head nil' do
      stack = Stack.new
      expect(stack.head).to be(nil)
    end

    it 'initializes with tail nil' do
      stack = Stack.new
      expect(stack.tail).to be(nil)
    end

    it 'initializes with length 0' do
      stack = Stack.new
      expect(stack.length).to eq(0)
    end
  end

  context 'push' do
    it 'adds items to the stack' do
      stack = n_item_stack(1)
      expect(stack.length).to eq(1)
    end

    it 'makes the first node the head and tail of the stack' do
      stack = n_item_stack(1)
      expect(stack.head).to eq(stack.tail).and have_attributes(data: 1)
    end

    it 'makes subsequent nodes the tail of the stack' do
      stack = n_item_stack(2)
      expect(stack.tail).to have_attributes(data: 2)
      expect(stack.tail).to_not eq(stack.head)
    end

    it 'sets the next attribute of the old tail to the new tail node' do
      stack = n_item_stack(2)
      old_tail = stack.tail
      stack.push(3)
      expect(old_tail.next).to eq(stack.tail)
    end
  end

  context 'pop' do
    it 'sets the tail of the stack to the old tail\'s preceding node' do
      stack = n_item_stack(3)
      stack.pop
      expect(stack.tail).to have_attributes(data: 2)
    end

    it 'reduces the stack\'s length correctly' do
      stack = n_item_stack(2)
      stack.pop
      expect(stack.length).to eq(1)
    end

    it 'returns the popped node' do
      stack = n_item_stack(1)
      old_tail = stack.tail
      expect(stack.pop).to eq(old_tail)
    end

    it 'returns nil if the stack is empty' do
      stack = Stack.new
      expect(stack.pop).to eq(nil)
    end
  end

  context 'peek' do
    it 'returns the current tail' do
      stack = n_item_stack(2)
      expect(stack.peek).to eq(stack.tail)
    end

    it 'does not modify the stack' do
      stack = n_item_stack(2)
      stack.peek
      expect(stack.length).to eq(2)
    end
  end

  context 'clear' do
    it 'removes each element from the stack' do
      stack = n_item_stack(5)
      stack.clear
      expect(stack.length).to eq(0)
    end
  end

  context 'each' do
    it 'iterates over each element in the stack' do
      stack = n_item_stack(3)
      stack_data = []
      stack.each { |node| stack_data << node.data }
      expect(stack_data).to eq([1, 2, 3])
    end

    it 'returns the unmodified stack' do
      stack = n_item_stack(3)
      expect(stack.each { |node| node.data * 2 }).to eq(stack)
    end

    it 'raises a LocalJumpError if no block is given' do
      stack = n_item_stack(3)
      expect{ stack.each }.to raise_error(LocalJumpError)
    end
  end

end