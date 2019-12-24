class Node
  attr_accessor :key, :data, :height, :left, :right, :deleted
  
  def initialize(key, data=nil)
    @key = key
    @data = data
    @height = 1
    @left = nil
    @right = nil
    @deleted = false
  end
end