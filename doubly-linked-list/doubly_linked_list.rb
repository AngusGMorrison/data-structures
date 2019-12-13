require_relative './node'
require_relative './doubly_linked_list_errors'

class DoublyLinkedList
  include DoublyLinkedListErrors

  attr_reader :head, :tail, :length

  def initialize
    @head = nil
    @tail = nil
    @length = 0
  end

  def insert(data)
    node = Node.new(data)
    if @head
      node.prev = @tail
      @tail.next = node
    else
      @head = node
    end
    @tail = node
    @length += 1
    node
  end

  def delete(node)
    # Ordinarily an O(1) operation, but the need to verify that the node is part
    # of the list makes it an O(n) operation
    raise NodeNotFound unless find { |existing_node| existing_node == node }
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
    preceding_node = node.prev
    next_node = node.next
    preceding_node.next = next_node
    next_node.prev = preceding_node if next_node
  end

  def concat(list)
    unless list.is_a?(DoublyLinkedList)
      raise ArgumentError.new("Expected a linked list, received #{list.class.name}")
    end
    @tail.next = list.head
    list.head.prev = @tail
    @tail = list.tail
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

  def find_last(&predicate)
    current = @tail
    while current
      return current if yield(current)
      current = current.prev
    end
  end

  def each(&block)
    current = @head
    while current
      yield(current)
      current = current.next
    end
  end

  def reverse_each(&block)
    current = @tail
    while current
      yield(current)
      current = current.prev
    end
  end

  def to_a
    array = []
    each { |node| array << node.data }
    array
  end

  def map(reverse: false, &block)
    mapped = DoublyLinkedList.new
    mapper = Proc.new do |node|
      value = block_given? ? yield(node.data) : node.data
      mapped.insert(value)
    end
    reverse ? reverse_each(&mapper) : each(&mapper)
    mapped
  end
end
