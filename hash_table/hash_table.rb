require_relative '../linked-list/linked_list'
require_relative './hash_table_errors'

class HashTable
  include HashTableErrors

  attr_reader :size, :array

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
    @array = Array.new(@bucket_count) { LinkedList.new }
    @size = 0
  end

  def insert(key, value)
    index = get_index_from_key(key)
    node = @array[index].find { |node| node.data[0] == key }
    if node
      node.data[1] = value
    elsif exceeds_max_load_factor?
      rehash
      return insert(key, value)
    else
      node = @array[index].insert([key, value])
      @size += 1
    end
    node.data[1]
  end

  # private def will_exceed_avg_bucket_size?
  #   (@size + 1) / @bucket_count.to_f > MAX_BUCKET_SIZE
  # end

  private def exceeds_max_load_factor?
    projected_load_factor > MAX_LOAD_FACTOR
  end

  private def projected_load_factor
    (@size + 1) / (@bucket_count * MAX_BUCKET_SIZE).to_f
  end

  private def rehash
    raise HashTableFull if @rehashes >= MAX_REHASHES
    @rehashes += 1
    @bucket_count = PRIMES[@rehashes]
    new_table = Array.new(@bucket_count) { LinkedList.new }
    @array.each do |list|
      list.each do |node|
        index = get_index_from_key(node.data[0])
        new_table[index].insert(node.data)
      end
    end
    @array = new_table
  end

  def get(key)
    index = get_index_from_key(key)
    return nil unless @array[index]
    node = find_node(@array[index], key)
    node ? node.data[1] : nil
  end

  def delete(key)
    index = get_index_from_key(key)
    return nil unless @array[index]
    node = find_node(@array[index], key)
    node ? delete_node(@array[index], node) : nil
  end

  private def find_node(list, key)
    list.find { |node| node.data[0] == key }
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
    @array.each do |list|
      list.each do |node|
        index ? extracted << node.data[index] : extracted << node.data
      end
    end
    extracted
  end

end