#  RBT properties
1. node's color is red or black
2. root node is black
3. all leaf node (nilnode) is black
4. if a node is red, both of its children is black
5. starting from any node, all simple paths dwon to leaf nodes hold the same number of black nodes.

# Derivation
## longest path, h(max) <= 2*shortest path, h(min)
### note. complete binary tree has n = 2^h+1 -1 ndoes.
> 2^h(min)+1 - 1 <= n <= 2^h(max)+1 -1
>> h(min) <= log(n+1) -1 <= h(max) <= 2*h(min)
>> h(min) = 𝞱(log n)

h(min) <=h(max) <=2*h(min) => h(max) = 𝞱(log n)


## Note. Complete Binary tree height:
- 對於一棵擁有  n  個節點的完全二元樹，其最大和最小高度分別為：
  - 最大高度： h_(max) = floor{ log_2(n) }
  - 最小高度： h_(min) = abs{ log_2(n + 1) }- 1
