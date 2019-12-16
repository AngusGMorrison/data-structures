require_relative './node'

class Queue
  attr_reader :head, :tail, :length

  def initialize
    @head = nil
    @tail = nil
    @length = 0
  end

  def push(data)
    # O(1)
    node = Node.new(data)
    @head ? @tail.next = node : @head = node
    @tail = node
    @length += 1
    node
  end

  def shift
    # O(1)
    return nil unless @head
    shifted_node = @head
    @head = @head.next
    @length -= 1
    @tail = nil if @length == 0
    shifted_node
  end

  def peek
    # O(1)
    @head
  end

  def clear
    # O(n)
    while @head
      shift 
    end
  end

  def each(&block)
    # O(n)
    current = @head
    while current
      yield(current)
      current = current.next
    end
    self
  end

end