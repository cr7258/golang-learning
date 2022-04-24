package main

import "testing"

/*
案例1：
• 王强第一次录入的时候，他的体脂是 38
• 王强第二次录入的时候，他的体脂是 32
• 这时，王强的最佳体脂是 32
• 李静录入他的体脂 28
• 李静的最佳体脂是 28
• 李静排名第一，体脂 28；王强排名第二，体脂 32。
*/
func TestCase1(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("王强", 0.32)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 1 {
			t.Fatalf("预期王强第一名, 但是得到的是: %d", randOfWQ)
			return
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率是 0.32, 但是得到的是: %f", fatRateOfWQ)
		}
	}
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二名, 但是得到的是: %d", randOfWQ)
			return
		}
		if fatRateOfWQ != 0.32 {
			t.Fatalf("预期王强的体脂率是 0.32, 但是得到的是: %f", fatRateOfWQ)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一名, 但是得到的是: %d", randOfLJ)
			return
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是 0.28, 但是得到的是: %f", fatRateOfLJ)
		}
	}
}

/*
案例2：
• 王强录入体脂是 38
• 张伟录入体脂是 38
• 李静录入体脂是 28
• 李静排名第一，体脂 28；王强、张伟排名第二，体脂 38。
*/
func TestCase2(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟", 0.38)
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二名, 但是得到的是: %d", randOfWQ)
			return
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂率是 0.38, 但是得到的是: %f", fatRateOfWQ)
		}
	}
	{
		randOfZW, fatRateOfZW := getRand("张伟")
		if randOfZW != 2 {
			t.Fatalf("预期张伟第二名, 但是得到的是: %d", randOfZW)
			return
		}
		if fatRateOfZW != 0.38 {
			t.Fatalf("预期张伟的体脂率是 0.38, 但是得到的是: %f", fatRateOfZW)
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一名, 但是得到的是: %d", randOfLJ)
			return
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是 0.28, 但是得到的是: %f", fatRateOfLJ)
		}
	}
}

/*
案例3：
• 王强录入体脂是 38
• 李静录入体脂是 28
• 张伟注册成功，未录入体脂
• 李静排名第一，体脂 28；王强排名第二，体脂 38；张伟排名第三，没有体脂记录。
*/
func TestCase3(t *testing.T) {
	inputRecord("王强", 0.38)
	inputRecord("张伟")
	inputRecord("李静", 0.28)
	{
		randOfWQ, fatRateOfWQ := getRand("王强")
		if randOfWQ != 2 {
			t.Fatalf("预期王强第二名, 但是得到的是: %d", randOfWQ)
			return
		}
		if fatRateOfWQ != 0.38 {
			t.Fatalf("预期王强的体脂率是 0.38, 但是得到的是: %f", fatRateOfWQ)
		}
	}
	{
		randOfZW, _ := getRand("张伟")
		if randOfZW != 3 {
			t.Fatalf("预期张伟第三名, 但是得到的是: %d", randOfZW)
			return
		}
	}
	{
		randOfLJ, fatRateOfLJ := getRand("李静")
		if randOfLJ != 1 {
			t.Fatalf("预期李静第一名, 但是得到的是: %d", randOfLJ)
			return
		}
		if fatRateOfLJ != 0.28 {
			t.Fatalf("预期李静的体脂率是 0.28, 但是得到的是: %f", fatRateOfLJ)
		}
	}
}
