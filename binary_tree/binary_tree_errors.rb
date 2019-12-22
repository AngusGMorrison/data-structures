module BinaryTreeErrors
  class NodeOverride < StandardError
    def message
      "Can't override an existing child node"
    end
  end

  class NilMerge < ArgumentError
    def message
      "Can't merge nil tree"
    end
  end
end