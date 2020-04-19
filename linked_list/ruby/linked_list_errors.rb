module LinkedListErrors
  class NodeNotFound < StandardError
    def message
      "Node not found"
    end
  end
end