package main

import ("reflect";"testing")

func TestInorderVariants(t *testing.T) {
    root:=&TreeNode{Val:1,Right:&TreeNode{Val:2,Left:&TreeNode{Val:3}}}
    bst:=&TreeNode{Val:2,Left:&TreeNode{Val:1},Right:&TreeNode{Val:3}}
    skew:=&TreeNode{Val:3,Left:&TreeNode{Val:2,Left:&TreeNode{Val:1}}}
    cases:=[]struct{name string; root *TreeNode; want []int}{
        {"nil",nil,[]int{}},
        {"single",&TreeNode{Val:7},[]int{7}},
        {"asymmetric",root,[]int{1,3,2}},
        {"BST",bst,[]int{1,2,3}},
        {"skew",skew,[]int{1,2,3}},
    }
    implementations:=[]struct{name string; f func(*TreeNode)[]int}{
        {"pointer",inorderTraversal},{"closure",inorderClosure},{"iterative",inorderIterative},
    }
    for _,tc:=range cases {
        for _,imp:=range implementations {
            t.Run(tc.name+"/"+imp.name,func(t *testing.T) {
                if got:=imp.f(tc.root);!reflect.DeepEqual(got,tc.want) {
                    t.Errorf("%s returned %v, want %v",imp.name,got,tc.want)
                }
            })
        }
    }
}
