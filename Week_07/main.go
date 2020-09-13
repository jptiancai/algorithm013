//208. 实现 Trie (前缀树)
// 把 v - 'a' 的差值 作为26个字母来处理

type Trie struct {
    next [26]*Trie
    isEnd bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this
    for _, v := range word {
        v -= 'a'
        if node.next[v] == nil {
            node.next[v] = &Trie{}
        }
        node = node.next[v]
    }
    node.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this
    for _, v := range word {
        if node = node.next[v - 'a']; node == nil {
            return false
        }
    }
    return node.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this
    for _, v := range prefix {
        if node = node.next[v - 'a']; node == nil {
            return false
        }
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */


 // 212. 单词搜索 II
// 1. dfs, 类似于岛屿问题
// 时间复杂度O(n^2)


type TrieNode struct {
	// 根据不同的题型具体变化，这里是匹配字符串，而不是找到，所以是 isEnd -> word
	word     string
	children [26]*TrieNode
}

func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	// insert trie
	for _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}

	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, board, root, &result)
		}
	}
	return result
}

func dfs(i, j int, board [][]byte, node *TrieNode, result *[]string) {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	// 访问过了或者字典中没有
	if c == '#' || node.children[c-'a'] == nil {
		return
	}
	node = node.children[c-'a']
	// 防止重复
	if node.word != "" {
		*result = append(*result, node.word)
		node.word = ""
	}

	// 访问过的地方设置 #
	board[i][j] = '#'
	dfs(i+1, j, board, node, result)
	dfs(i-1, j, board, node, result)
	dfs(i, j+1, board, node, result)
	dfs(i, j-1, board, node, result)
	// reverse state
	board[i][j] = c
}


// dfs 的解法

func findWords(board [][]byte, words []string) []string {
    trie := Trie{}
    for _, word := range words {
        trie.Insert(word)
    }

    res := []string{}

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            trie.Search(&board, i, j, &res)
        }
    }

    return res
}

// Trie - Prefix Tree
type Trie struct {
    links [26]*Trie
    word  string
}

// IsLeafNode - check if it is a leaf node
func (tr *Trie) IsLeafNode() bool {
    for _, link := range tr.links {
        if link != nil {
            return false
        }
    }
    return true
}

// Insert - insert a word
func (tr *Trie) Insert(word string) {
    for i := 0; i < len(word); i++ {
        ch := word[i] - 'a'
        if tr.links[ch] == nil {
            tr.links[ch] = &Trie{}
        }
        tr = tr.links[ch]
    }
    tr.word = word
}

// Search - search words
func (tr *Trie) Search(board *[][]byte, x, y int, res *[]string) {
    if x < 0 || x >= len(*board) || y < 0 || y >= len((*board)[0]) ||
        (*board)[x][y] == '#' || tr.links[(*board)[x][y]-'a'] == nil {
        return
    }

    letter := (*board)[x][y]
    cur := tr.links[letter-'a']

    // check if there is any match
    if cur.word != "" {
        (*res) = append(*res, cur.word)
        cur.word = ""
    }

    (*board)[x][y] = '#' // 标记已访问
    dx, dy := [4]int{1, 0, -1, 0}, [4]int{0, 1, 0, -1}
    for i := 0; i < 4; i++ {
        cur.Search(board, x+dx[i], y+dy[i], res)
    }
    (*board)[x][y] = letter // restore the original letter

    // 使用剪枝进行优化
    // 对于 Trie 中的叶节点，一旦遍历它（即找到匹配的单词），就不需要再遍历它了。
    if cur.IsLeafNode() {
        tr.links[letter-'a'] = nil
    }
}