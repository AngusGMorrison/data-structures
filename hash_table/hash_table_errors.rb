module HashTableErrors

  class HashTableFull < StandardError
    def message
      'Hash table is full'
    end
  end
end