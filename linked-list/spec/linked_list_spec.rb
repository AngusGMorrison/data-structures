require 'rspec'
require_relative '../linked_list'

RSpec.describe LinkedList do

  it 'is initialized with length 0' do
    list = LinkedList.new
    expect(list.length).to eq(0)
  end

  it 'is initialized with head nil' do
    list = LinkedList.new
    expect(list.head).to be(nil)
  end

  it 'is initialized with tail nil' do
    list = LinkedList.new
    expect(list.tail).to be(nil)
  end

  it 'allows a new node to become the head and tail of an empty list' do
    list = LinkedList.new
    list.insert(1)
    expect(list.head).to eq(list.tail).and have_attributes(data: 1)
  end

  it 'allows new nodes to be added to the end of a list' do
    # list = LinkedList.new
    # list.insert(1)
    # list.insert(2)
    # expect(list.tail).to eq(node2)
  end

  # it 'allows body nodes to be deleted from the list' do
  #   list.insert
  # end


end