"""
    Very basic implementation
"""

class TrieNode():
    
    
    def __init__(self, char: None):
        self.char = char
        self.children = []

class Trie:
    def __init__(self):
        self.root = TrieNode('*')
    
    def add_word(self, word:str):
        node = self.root
        for char in word:
            found_in_child = False
            for child in node.children:
                if child.char == char:
                    found_in_child = True
                    node = child
                    break
            if not found_in_child:
                t = TrieNode(char)
                node.children.append(t)
                node = t
    
    def find(self, word:str):
        node = self.root
        counter = 0
        for char in word:
            found_in_child = False
            for child in node.children:
                if child.char == char:
                    counter += 1
                    found_in_child = True
                    node = child
                    break
            if not found_in_child:
                return False, counter
        return True, counter

if __name__ == '__main__':
    t = Trie()
    t.add_word("tree")
    t.add_word("troo")
    t.add_word("fix")
    assert t.find("tr") == (True, 2)
    assert t.find("treee") == (False, 4)
    assert t.find("troo") == (True, 4)
    assert len(t.root.children) == 2 
