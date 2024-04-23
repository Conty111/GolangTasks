package ozon_techpoint

import (
	"bufio"
	"fmt"
	"os"
)

func Train2() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	text, _ := in.ReadBytes('\n')
	text = text[:len(text)-1]
	var start, end, n, i, j int16
	var nakl string
	fmt.Fscan(in, &n)
	for i = 0; i < n; i++ {
		fmt.Fscanf(in, "%d %d %s", &start, &end, &nakl)
		fmt.Println(nakl)
		for j = 0; j < start+end+1; j++ {
			text[j+start+1] = nakl[j]
		}
		fmt.Println(text)
	}
	fmt.Fprintf(out, "%s", text)
}
