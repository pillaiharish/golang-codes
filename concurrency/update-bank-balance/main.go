package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

var wg sync.WaitGroup

func main(){
	// variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting value
	fmt.Printf("Starting bank balance is $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Slavery", Amount: 500},
		{Source: "Medium Blogging", Amount: 2},
		{Source: "Teaching kids", Amount: 33},
		{Source: "Consultant", Amount: 1000},
	}

	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is made;
	// keep a running total
	for i, income := range(incomes) {
		go func(i int, income Income) {
			defer wg.Done()

			for week:=1 ; week<=52; week++ {

				balance.Lock()	
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d Updated bank balance for source %s is $%d\n", week, income.Source, bankBalance)
			}
		}(i, income)
	}
	wg.Wait()

	// print out final balance
	fmt.Printf("Final bank balance is $%d.00", bankBalance)
	fmt.Println()
}




// Output:
/*
Starting bank balance is $0.00
Updated bank balance for source Consultant is $1000
Updated bank balance for source Consultant is $2535
Updated bank balance for source Consultant is $3535
Updated bank balance for source Consultant is $4535
Updated bank balance for source Consultant is $5535
Updated bank balance for source Consultant is $6535
Updated bank balance for source Consultant is $7535
Updated bank balance for source Consultant is $8535
Updated bank balance for source Consultant is $9535
Updated bank balance for source Consultant is $10535
Updated bank balance for source Consultant is $11535
Updated bank balance for source Consultant is $12535
Updated bank balance for source Consultant is $13535
Updated bank balance for source Consultant is $14535
Updated bank balance for source Slavery is $1535
Updated bank balance for source Slavery is $16035
Updated bank balance for source Slavery is $16535
Updated bank balance for source Slavery is $17035
Updated bank balance for source Slavery is $17535
Updated bank balance for source Slavery is $18035
Updated bank balance for source Slavery is $18535
Updated bank balance for source Slavery is $19035
Updated bank balance for source Slavery is $19535
Updated bank balance for source Slavery is $20035
Updated bank balance for source Slavery is $20535
Updated bank balance for source Slavery is $21035
Updated bank balance for source Slavery is $21535
Updated bank balance for source Slavery is $22035
Updated bank balance for source Consultant is $15535
Updated bank balance for source Consultant is $23535
Updated bank balance for source Consultant is $24535
Updated bank balance for source Consultant is $25535
Updated bank balance for source Consultant is $26535
Updated bank balance for source Consultant is $27535
Updated bank balance for source Consultant is $28535
Updated bank balance for source Consultant is $29535
Updated bank balance for source Consultant is $30535
Updated bank balance for source Consultant is $31535
Updated bank balance for source Slavery is $22535
Updated bank balance for source Slavery is $33035
Updated bank balance for source Slavery is $33535
Updated bank balance for source Slavery is $34035
Updated bank balance for source Slavery is $34535
Updated bank balance for source Consultant is $32535
Updated bank balance for source Consultant is $36035
Updated bank balance for source Consultant is $37035
Updated bank balance for source Consultant is $38035
Updated bank balance for source Consultant is $39035
Updated bank balance for source Consultant is $40035
Updated bank balance for source Consultant is $41035
Updated bank balance for source Consultant is $42035
Updated bank balance for source Consultant is $43035
Updated bank balance for source Consultant is $44035
Updated bank balance for source Consultant is $45035
Updated bank balance for source Consultant is $46035
Updated bank balance for source Consultant is $47035
Updated bank balance for source Consultant is $48035
Updated bank balance for source Consultant is $49035
Updated bank balance for source Consultant is $50035
Updated bank balance for source Consultant is $51035
Updated bank balance for source Consultant is $52035
Updated bank balance for source Consultant is $53035
Updated bank balance for source Consultant is $54035
Updated bank balance for source Consultant is $55035
Updated bank balance for source Consultant is $56035
Updated bank balance for source Consultant is $57035
Updated bank balance for source Consultant is $58035
Updated bank balance for source Consultant is $59035
Updated bank balance for source Consultant is $60035
Updated bank balance for source Consultant is $61035
Updated bank balance for source Consultant is $62035
Updated bank balance for source Medium Blogging is $1002
Updated bank balance for source Slavery is $35035
Updated bank balance for source Slavery is $62537
Updated bank balance for source Slavery is $63037
Updated bank balance for source Slavery is $63537
Updated bank balance for source Slavery is $64037
Updated bank balance for source Slavery is $64537
Updated bank balance for source Slavery is $65037
Updated bank balance for source Slavery is $65537
Updated bank balance for source Slavery is $66037
Updated bank balance for source Slavery is $66537
Updated bank balance for source Slavery is $67037
Updated bank balance for source Slavery is $67537
Updated bank balance for source Slavery is $68037
Updated bank balance for source Slavery is $68537
Updated bank balance for source Slavery is $69037
Updated bank balance for source Slavery is $69537
Updated bank balance for source Slavery is $70037
Updated bank balance for source Medium Blogging is $62037
Updated bank balance for source Medium Blogging is $70539
Updated bank balance for source Medium Blogging is $70541
Updated bank balance for source Medium Blogging is $70543
Updated bank balance for source Medium Blogging is $70545
Updated bank balance for source Medium Blogging is $70547
Updated bank balance for source Medium Blogging is $70549
Updated bank balance for source Medium Blogging is $70551
Updated bank balance for source Medium Blogging is $70553
Updated bank balance for source Medium Blogging is $70555
Updated bank balance for source Medium Blogging is $70557
Updated bank balance for source Medium Blogging is $70559
Updated bank balance for source Medium Blogging is $70561
Updated bank balance for source Medium Blogging is $70563
Updated bank balance for source Medium Blogging is $70565
Updated bank balance for source Medium Blogging is $70567
Updated bank balance for source Medium Blogging is $70569
Updated bank balance for source Medium Blogging is $70571
Updated bank balance for source Medium Blogging is $70573
Updated bank balance for source Medium Blogging is $70575
Updated bank balance for source Medium Blogging is $70577
Updated bank balance for source Medium Blogging is $70579
Updated bank balance for source Medium Blogging is $70581
Updated bank balance for source Medium Blogging is $70583
Updated bank balance for source Medium Blogging is $70585
Updated bank balance for source Medium Blogging is $70587
Updated bank balance for source Medium Blogging is $70589
Updated bank balance for source Medium Blogging is $70591
Updated bank balance for source Medium Blogging is $70593
Updated bank balance for source Medium Blogging is $70595
Updated bank balance for source Medium Blogging is $70597
Updated bank balance for source Medium Blogging is $70599
Updated bank balance for source Medium Blogging is $70601
Updated bank balance for source Medium Blogging is $70603
Updated bank balance for source Medium Blogging is $70605
Updated bank balance for source Medium Blogging is $70607
Updated bank balance for source Medium Blogging is $70609
Updated bank balance for source Medium Blogging is $70611
Updated bank balance for source Medium Blogging is $70613
Updated bank balance for source Medium Blogging is $70615
Updated bank balance for source Medium Blogging is $70617
Updated bank balance for source Medium Blogging is $70619
Updated bank balance for source Medium Blogging is $70621
Updated bank balance for source Medium Blogging is $70623
Updated bank balance for source Medium Blogging is $70625
Updated bank balance for source Medium Blogging is $70627
Updated bank balance for source Medium Blogging is $70629
Updated bank balance for source Medium Blogging is $70631
Updated bank balance for source Medium Blogging is $70633
Updated bank balance for source Medium Blogging is $70635
Updated bank balance for source Medium Blogging is $70637
Updated bank balance for source Slavery is $70537
Updated bank balance for source Teaching kids is $1035
Updated bank balance for source Slavery is $71170
Updated bank balance for source Slavery is $71670
Updated bank balance for source Slavery is $72170
Updated bank balance for source Slavery is $72670
Updated bank balance for source Slavery is $73170
Updated bank balance for source Slavery is $73670
Updated bank balance for source Slavery is $74170
Updated bank balance for source Slavery is $74670
Updated bank balance for source Slavery is $75170
Updated bank balance for source Slavery is $75670
Updated bank balance for source Slavery is $76170
Updated bank balance for source Slavery is $76670
Updated bank balance for source Slavery is $77170
Updated bank balance for source Slavery is $77670
Updated bank balance for source Slavery is $78170
Updated bank balance for source Teaching kids is $70670
Updated bank balance for source Teaching kids is $78203
Updated bank balance for source Teaching kids is $78236
Updated bank balance for source Teaching kids is $78269
Updated bank balance for source Teaching kids is $78302
Updated bank balance for source Teaching kids is $78335
Updated bank balance for source Teaching kids is $78368
Updated bank balance for source Teaching kids is $78401
Updated bank balance for source Teaching kids is $78434
Updated bank balance for source Teaching kids is $78467
Updated bank balance for source Teaching kids is $78500
Updated bank balance for source Teaching kids is $78533
Updated bank balance for source Teaching kids is $78566
Updated bank balance for source Teaching kids is $78599
Updated bank balance for source Teaching kids is $78632
Updated bank balance for source Teaching kids is $78665
Updated bank balance for source Teaching kids is $78698
Updated bank balance for source Teaching kids is $78731
Updated bank balance for source Teaching kids is $78764
Updated bank balance for source Teaching kids is $78797
Updated bank balance for source Teaching kids is $78830
Updated bank balance for source Teaching kids is $78863
Updated bank balance for source Teaching kids is $78896
Updated bank balance for source Teaching kids is $78929
Updated bank balance for source Teaching kids is $78962
Updated bank balance for source Teaching kids is $78995
Updated bank balance for source Teaching kids is $79028
Updated bank balance for source Teaching kids is $79061
Updated bank balance for source Teaching kids is $79094
Updated bank balance for source Teaching kids is $79127
Updated bank balance for source Teaching kids is $79160
Updated bank balance for source Teaching kids is $79193
Updated bank balance for source Teaching kids is $79226
Updated bank balance for source Teaching kids is $79259
Updated bank balance for source Teaching kids is $79292
Updated bank balance for source Teaching kids is $79325
Updated bank balance for source Teaching kids is $79358
Updated bank balance for source Teaching kids is $79391
Updated bank balance for source Teaching kids is $79424
Updated bank balance for source Teaching kids is $79457
Updated bank balance for source Teaching kids is $79490
Updated bank balance for source Teaching kids is $79523
Updated bank balance for source Teaching kids is $79556
Updated bank balance for source Teaching kids is $79589
Updated bank balance for source Teaching kids is $79622
Updated bank balance for source Teaching kids is $79655
Updated bank balance for source Teaching kids is $79688
Updated bank balance for source Teaching kids is $79721
Updated bank balance for source Teaching kids is $79754
Updated bank balance for source Teaching kids is $79787
Updated bank balance for source Teaching kids is $79820
Final bank balance is $79820.00
*/
