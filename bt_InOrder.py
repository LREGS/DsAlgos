def walk (node: int, path: list): 
    if not node:
        return path
    
    path.append(node.value)

    walk(node.left, path)
    walk(node.right, path)

    return path 

      n 

