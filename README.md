## Run Locally

Clone the project

```bash
  https://github.com/sagarsaud51/english-word-validation-api
```

Go to the project directory

```bash
  cd english-word-validation-api
```

Start the server

```bash
  go run src/cmd/main.go
```

The application should now be running and accessible at [http://localhost:8080](http://localhost:8080).


## Code Explanation

### TrieNode Struct

```go
type TrieNode struct {
    children map[rune]*TrieNode
    isEnd    bool
}
```
* `TrieNode` represents a node in the Trie data structure.
* It contains a map of child nodes (`children`) and a flag (`isEnd`) indicating whether the node marks the end of a valid word.

### Trie Struct
```go
type Trie struct {
    root *TrieNode
}
```
* `Trie` is the main structure that holds the Trie data structure.
* It contains a pointer to the `root` node.

### NewTrie()
```go
func NewTrie() *Trie {
    return &Trie{
        root: &TrieNode{
            children: make(map[rune]*TrieNode),
        },
    }
}
```
* `NewTrie()` initializes and returns a new Trie.
* It creates a new Trie with an empty root node and initializes the children map.

### Insert(word)
```go
func (t *Trie) Insert(word string) {
    node := t.root
    for _, char := range word {
        if node.children[char] == nil {
            node.children[char] = &TrieNode{
                children: make(map[rune]*TrieNode),
            }
        }
        node = node.children[char]
    }
    node.isEnd = true
}
```
* `Insert` adds a word to the Trie.
* It iterates over each character in the word and checks if a child node for that character exists. If not, it creates a new node.
* It then moves to the child node and repeats the process.
* Finally, it marks the last node as the end of a valid word.


### Search(word)
```go
func (t *Trie) Search(word string) bool {
    node := t.root
    for _, char := range word {
        if node.children[char] == nil {
            return false
        }
        node = node.children[char]
    }
    return node.isEnd
}

```
* `Search` checks if a given word exists in the Trie.
* It iterates over each character in the word, following the Trie structure.
* If at any point a child node for a character is missing, it means the word is not in the Trie, and false is returned.
* If the loop completes without issue, it means the entire word is present in the Trie, and true is returned.


## Visual Representation of Trie

```JSON
{
  "root": {
    "children": {
      "a": {
        "children": {
          "\r": {
            "children": {},
            "isEnd": true
          },
          "a": {
            "children": {
              "\r": {
                "children": {},
                "isEnd": true
              },
              "a": {
                "children": {
                  "\r": {
                    "children": {},
                    "isEnd": true
                  }
                },
                "isEnd": false
              },
              "h": {
                "children": {
                  "\r": {
                    "children": {},
                    "isEnd": true
                  },
                  "e": {
                    "children": {
                      "d": {
                        "children": {
                          "\r": {
                            "children": {},
                            "isEnd": true
                          }
                        },
                        "isEnd": false
                      }
                    },
                    "isEnd": false
                  },
                  "i": {
                    "children": {
                      "n": {
                        "children": {
                          "g": {
                            "children": {
                              "\r": {
                                "children": {},
                                "isEnd": true
                              }
                            },
                            "isEnd": false
                          }
                        },
                        "isEnd": false
                      }
                    },
                    "isEnd": false
                  }
                },
                "isEnd": false
              }
            },
            "isEnd": false
          }
        },
        "isEnd": false
      },
      "z": {
        "children": {},
        "isEnd": true
      }
    },
    "isEnd": false
  }
}

```