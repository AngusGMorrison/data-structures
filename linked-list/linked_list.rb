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
    node = Node.new(data)
    if @head
      @tail.next = node
    else
      @head = node
    end
    @tail = node
    @length += 1
  end

  def delete(node)
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
    while preceding node && preceding_node.next != node
      preceding_node = preceding_node.next
    end
    raise NodeNotFound unless preceding_node.next
    preceding_node.next = node.next
  end

  def concat(list)
    unless list.is_a?(LinkedList)
      raise ArgumentError("Expected a linked list, received #{list.class.name}")
    end
    @tail.next = list.head
    @length += list.length
  end

  def clear
    while @length > 0
      delete(head)
    end
  end

  def find(&predicate)
    current = @head
    while current
      return current if yield(current)
      current = current.next
    end
  end

  def each(&predicate)
    current = @head
    while current
      yield(current)
      current = current.next
    end
  end

  def print
    each { |node| puts node.data }
  end
end
