<script>
	function add(a , b)
	{
		// for loop will start from 1 and move till the value of second
		// number , first number(a) is incremented in for loop
		for (i = 1; i <= b; i++)
		{
			a++;
		}
		return a;
	}

	// first number is 10 and second number is 32 , for loop will start
	var a = add(10, 32);
	
	// from 1 and move till 32 and the value of a is incremented 32 times
	// which will give us the total sum of two numbers
	document.write(a);

// This code is contributed by Rajput-Ji
</script>
