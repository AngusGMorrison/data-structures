require_relative './node'

class Stack
  attr_reader :head, :tail, :length

  def initialize
    @head = nil
    @tail = nil
    @length = 0
  end

  def push(data)
    new_node = Node.new(data)
    if @head
      @tail.next = new_node
      @tail = new_node
    else
      @head = @tail = new_node
    end
    @length += 1
    new_node
  end

  def pop
    return nil unless @tail
    node_to_pop = @tail
    pop_node(node_to_pop)
    node_to_pop
  end

  private def pop_node(node_to_pop)
    if node_to_pop == @head
      @head = @tail = nil
    else
      @tail = get_new_tail(node_to_pop)
      @tail.next = nil
    end
    @length -= 1
  end

  private def get_new_tail(node_to_pop)
    new_tail = @head
    while new_tail.next != node_to_pop
      new_tail = new_tail.next
    end
    new_tail
  end

  def peek
    @tail
  end

  def clear
    while @head
      pop
    end
  end

  def each(&block)
    current = @head
    while current
      yield(current)
      current = current.next
    end
    self
  end
end