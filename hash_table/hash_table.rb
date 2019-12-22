require_relative '../linked-list/linked_list'
require_relative './hash_table_errors'

class HashTable
  include HashTableErrors

  attr_reader :size, :table

  PRIMES = [
    8 + 3,
    16 + 3,
    32 + 5,
    64 + 3,
    128 + 3,
    256 + 27,
    512 + 9
  ]
  MAX_REHASHES = PRIMES.length
  MAX_BUCKET_SIZE = 5
  MAX_LOAD_FACTOR = 0.5
  

  def initialize
    @rehashes = 0
    @bucket_count = PRIMES[@rehashes]
    @table = Array.new(@bucket_count) { LinkedList.new }
    @size = 0
  end

  def insert(key, value)
    # ~O(1)
    node = find_node(key)
    node ? update_node(node, value) : create_node(key, value)
  end

  private def update_node(node, value)
    node.data[1] = value
  end

  private def create_node(key, value)
    rehash if will_exceed_max_load_factor?
    index = get_index_from_key(key)
    node = @table[index].insert([key, value])
    @size += 1
    node.data[1]
  end

  private def will_exceed_max_load_factor?
    projected_load_factor > MAX_LOAD_FACTOR
  end

  private def projected_load_factor
    (@size + 1) / (@bucket_count * MAX_BUCKET_SIZE).to_f
  end

  private def rehash
    # O(n)
    raise HashTableFull if @rehashes >= MAX_REHASHES
    @rehashes += 1
    @bucket_count = PRIMES[@rehashes]
    @table = upscale_table
  end

  private def upscale_table
    new_table = Array.new(@bucket_count) { LinkedList.new }
    @table.each do |list|
      list.each do |node|
        index = get_index_from_key(node.data[0])
        new_table[index].insert(node.data)
      end
    end
    new_table
  end

  def get(key)
    # ~O(1)
    node = find_node(key)
    node ? node.data[1] : nil
  end

  def delete(key)
    # ~O(1)
    node = find_node(key)
    index = get_index_from_key(key)
    node ? delete_node(@table[index], node) : nil  
  end

  private def find_node(key)
    index = get_index_from_key(key)
    @table[index].find { |node| node.data[0] == key }
  end

  private def delete_node(list, node)
    list.delete(node)
    @size -= 1
    node.data[1]
  end

  private def get_index_from_key(key)
    key_hash = hash_key(key)
    key_hash % @bucket_count
  end

  private def hash_key(key)
    key.to_s.hash
  end

  def entries
    # O(n)
    return extract
  end

  def keys
    return extract(index: 0)
  end

  def values
    return extract(index: 1)
  end

  private def extract(index: nil)
    extracted = []
    @table.each do |list|
      list.each do |node|
        index ? extracted << node.data[index] : extracted << node.data
      end
    end
    extracted
  end

end