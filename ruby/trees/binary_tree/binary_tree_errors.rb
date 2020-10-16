module BinaryTreeErrors
  class NodeOverride < StandardError
    def message
      "Can't override an existing child node"
    end
  end
end