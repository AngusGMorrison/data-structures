module LinkedListErrors
  class NodeNotFound < StandardError
    def message
      "Node not found"
    end
  end

  class BlockNotGiven < ArgumentError
    def message
      "Expected to receive a block"
    end
  end
end