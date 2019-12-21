require 'rspec'
require_relative '../hash_table'

RSpec.describe HashTable do

  test_key = :test
  test_value = 'pass'

  let(:create_table) do 
    table = HashTable.new
    table.insert(test_key, test_value)
    table
  end

  def n_item_table(n)
    table = HashTable.new
    n.times do |n|
      m = n + 1
      table.insert(m.to_s, m)
    end
    table
  end

  context 'initialize' do
    it 'initializes a hash table with size 0' do
      table = HashTable.new
      expect(table.size).to eq(0)
    end
  end

  context 'insert/get' do
    size_before_first_rehash = (HashTable::PRIMES[0] * HashTable::MAX_BUCKET_SIZE) / 2

    it 'inserts a key-value pair into the table' do
      table = create_table
      expect(table.get(test_key)).to eq(test_value)
    end

    it 'increases the table\'s size attribute when a key-value pair is inserted' do
      table = create_table
      expect(table.size).to eq(1)
    end

    it 'updates the value if the key already exists' do
      table = create_table
      new_test_value = 'new_value'
      table.insert(test_key, new_test_value)
      expect(table.get(test_key)).to eq(new_test_value)
      expect(table.size).to eq(1)
    end

    it 'rehashes the table when max load factor is exceeded' do
      hash = n_item_table(size_before_first_rehash)
      expect(hash.table.length).to eq(HashTable::PRIMES[0])
      hash.insert(test_key, test_value)
      expect(hash.table.length).to eq(HashTable::PRIMES[1])
    end

    it 'retrieves a value for a given key after a rehash has taken place' do
      hash = n_item_table(size_before_first_rehash)
      expect(hash.get('1')).to eq(1)
      hash = n_item_table(size_before_first_rehash + 1)
      p hash.entries
      expect(hash.get('1')).to eq(1)
    end
  end

  context 'delete' do
    it 'deletes an entry from the table given a matching key' do
      table = n_item_table(2)
      table.delete('1')
      expect(table.get('1')).to eq(nil)
    end

    it 'returns the value of the deleted entry' do
      table = n_item_table(2)
      expect(table.delete('1')).to eq(1)
    end

    it 'returns nil if the entry is not found' do
      table = HashTable.new
      expect(table.delete(test_key)).to eq(nil)
    end
  end

  context 'entries' do
    it 'returns an array of all key-value pairs' do
      table = n_item_table(2)
      expect(table.entries).to include(['1', 1], ['2', 2])
    end

    it 'returns an empty array if the hash table is empty' do
      table = HashTable.new
      expect(table.entries).to eq([])
    end
  end

  context 'keys' do
    it 'returns an array of all keys' do
      table = n_item_table(2)
      expect(table.keys).to include('1', '2')
    end

    it 'returns an empty array if the hash table is empty' do
      table = HashTable.new
      expect(table.keys).to eq([])
    end
  end

  context 'values' do
    it 'returns an array of all values' do
      table = n_item_table(2)
      expect(table.values).to include(1, 2)
    end

    it 'returns an empty array if the hash table is empty' do
      table = HashTable.new
      expect(table.keys).to eq([])
    end
  end

end