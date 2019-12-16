require 'rspec'
require_relative '../queue'

RSpec.describe Queue do

  def n_item_queue(n)
    queue = Queue.new
    for i in 1..n
      queue.push(i)
    end
    queue
  end

  context 'initialize' do
    queue = Queue.new

    it 'initializes with head nil' do
      expect(queue.head).to be(nil)
    end

    it 'initializes with tail nil' do
      expect(queue.tail).to be(nil)
    end

    it 'initializes with length 0' do
      expect(queue.length).to eq(0)
    end
  end

  context 'push' do
    it 'adds items to the queue' do
      queue = n_item_queue(1)
      expect(queue.length).to eq(1)
    end

    it 'makes the first item added the head and tail of the queue' do
      queue = n_item_queue(1)
      expect(queue.head).to eq(queue.tail).and have_attributes(data: 1)
    end

    it 'makes subsequent items the tail of the queue' do
      queue = n_item_queue(3)
      expect(queue.tail).to have_attributes(data: 3)
    end
  end

  context 'shift' do
    it 'removes items from the queue' do
      queue = n_item_queue(3)
      queue.shift
      expect(queue.length).to eq(2)
    end

    it 'removes the head of the queue' do
      queue = n_item_queue(3)
      original_head = queue.head
      expect(queue.shift).to eq(original_head)
      expect(queue.head).to eq(original_head.next)
    end

    it 'returns nil if queue is empty' do
      queue = Queue.new
      expect(queue.shift).to eq(nil)
    end

    it 'sets head and tail to nil if final node is removed' do
      queue = n_item_queue(1)
      queue.shift
      expect(queue.head).to be(nil).and eq(queue.tail)
    end
  end

  context 'peek' do
    it 'returns the current head node' do
      queue = n_item_queue(1)
      expect(queue.peek).to eq(queue.head)
    end

    it 'returns nil if the list is empty' do
      queue = Queue.new
      expect(queue.peek).to be(nil)
    end
  end

  context 'clear' do
    it 'deletes every node from the list' do
      queue = n_item_queue(10)
      queue.clear
      expect(queue.length).to eq(0)
      expect(queue.head).to eq(queue.tail).and be(nil)
    end
  end

  context 'each' do
    it 'iterates over each node in the list' do
      queue = n_item_queue(3)
      output = []
      queue.each { |node| output << node.data }
      expect(output).to eq([1, 2, 3])
    end

    it 'returns the original queue' do
      queue = n_item_queue(3)
      expect(queue.each { |node| node }).to eq(queue)
    end

    it 'raises a LocalJumpError if no block is given' do
      queue = n_item_queue(1)
      expect{ queue.each }.to raise_error(LocalJumpError)
    end
  end


end