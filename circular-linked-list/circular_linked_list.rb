require_relative './node'
require_relative './circular_linked_list_errors'


class CircularLinkedList
  include CircularLinkedListErrors

  attr_reader :head, :current, :length

  def initialize
    @head = nil
    @current = nil
    @length = 0
  end

  def insert(data)
    # O(n)
    if @length == 0
      insert_next(nil, data)
    elsif @length == 1
      insert_next(@head, data)
    else
      insert_next(last_node, data)
    end
  end

  private def insert_next(prev_node, data)
    new_node = Node.new(data)
    if !prev_node
      @head = new_node.next = new_node
    else
      new_node.next = prev_node.next
      prev_node.next = new_node
    end
    @length += 1
    new_node 
  end

  def last_node
    @current = @head
    for i in 0...(@length - 1)
      move_next
    end
    return @current
  end

  private def move_next
    @current = @current.next
  end

  def delete(node)
    # O(n)
    raise NodeNotFound unless find { |existing_node| existing_node == node }  
    preceding_node = get_preceding_node(node)
    return update_preceding_node(preceding_node)
  end

  private def get_preceding_node(node)
    preceding_node = @current = @head
    while move_next != node
      preceding_node = @current
    end
    return preceding_node
  end

  private def update_preceding_node(preceding_node)
    node_to_delete = preceding_node.next
    if node_to_delete == preceding_node
      @head = nil
    elsif @head == node_to_delete
      @head = node_to_delete.next
      preceding_node.next = @head
    else
      preceding_node.next = node_to_delete.next
    end

    @length -= 1
  end

  def clear
    # O(n)
    while @length > 0
      delete(@head)
    end
  end

  def concat(input_list)
    unless input_list.is_a?(CircularLinkedList)
      raise ArgumentError.new("Expected a CircularLinkedList, received #{input_list.class.name}")
    end
    # p last_node
    last_node.next = input_list.head
    last_input_node = input_list.last_node
    last_input_node.next = @head
    @length += input_list.length
    self
  end

  def each(&block)
    # O(n)
    @current = @head
    loop do
      yield(@current)
      break if move_next == @head
    end
    self
  end

  def to_a
    # O(n)
    array = []
    each { |node| array << node.data }
    array
  end

  def map(&block)
    new_list = CircularLinkedList.new
    mapper = nil
    if block_given?
     mapper = Proc.new { |node| new_list.insert(yield(node.data)) }
    else
      mapper = Proc.new { |node| new_list.insert(node.data) }
    end
    each(&mapper)
    new_list
  end

  def find(&predicate)
    # O(n)
    @current = @head
    loop do
      return @current if yield(@current)
      return nil if move_next == @head
    end
  end

  

  end

return