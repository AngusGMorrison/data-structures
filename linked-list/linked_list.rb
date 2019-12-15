require_relative './node'
require_relative './linked_list_errors'

class LinkedList
  include LinkedListErrors

  attr_reader :head, :tail, :length

  def initialize
    @head = nil
    @tail = nil
    @length = 0
  end

  def insert(data)
    #O(1)
    node = Node.new(data)
    @head ? @tail.next = node : @head = node
    @tail = node
    @length += 1
    node
  end

  def delete(node)
    #O(n)
    node == @head ? delete_head : delete_from_body(node)
    @length -= 1
  rescue NodeNotFound => message
    puts message
  end

  private def delete_head
    if @head.next
      @head = @head.next
    else
      @head = @tail = nil
    end 
  end

  private def delete_from_body(node)
    preceding_node = @head
    while preceding_node && preceding_node.next != node
      preceding_node = preceding_node.next
    end
    raise NodeNotFound unless preceding_node.next
    preceding_node.next = node.next
  end

  def concat(list)
    #O(1)
    unless list.is_a?(LinkedList)
      raise ArgumentError.new("Expected a linked list, received #{list.class.name}")
    end
    @tail.next = list.head
    @tail = list.tail
    @length += list.length
  end

  def clear
    #O(n)
    while @length > 0
      delete(head)
    end
  end

  def find(&predicate)
    #O(n)
    current = @head
    while current
      return current if yield(current)
      current = current.next
    end
  end

  def each(&block)
    #O(n)
    current = @head
    while current
      yield(current)
      current = current.next
    end
    self
  end

  def to_a
    #O(n)
    array = []
    each { |node| array << node.data }
    array
  end

  def map(&block)
    #O(n)
    mapped = LinkedList.new
    each do |node|
      value = block_given? ? yield(node.data) : node.data
      mapped.insert(value)
    end
    mapped
  end
end
