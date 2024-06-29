package main

import("fmt")
func findCenter(edges [][]int) int {
    edge1 := edges[0]
    edge2 := edges[1]

    if edge1[0] == edge2[0]{
        return edge1[0]
    }else if edge1[0] == edge2[1]{
        return edge1[0]
    }else{
        return edge1[1]
    }
    return 0
}


func main(){
	graph := [][]int{{1, 2}, {2, 3}, {4, 2}}
	center := findCenter(graph)
	fmt.Println(center)
}