## 贪心算法
- 场景： 设计一个公平的算法，实现从一堆牌中取出指定数目的牌

- 关键词：  公平（1/n） 、随机（rand）

- 解析

    假定这堆牌是一个数组 arr,长度为 n , 要实现公平即每次为1/n , 取一定数目的话，我们可以通过把arr 打乱然后截取指定长度的切片
    ```
        for i:=len(arr)-1;i?0;i--{
            swap(arr[i],rand(arr,i))
        }
        
    ```
     假设我们的arr有5个元素 1，2，3，4，5 ； 
     我们每次取最后一个和前面的任意一个做交换，此时最后一个数交换后得到的数字出现在最后一个位置的概率是 1/5
     然后排除最后一个，取倒数第二个和之前的任意一个做交换，此时交换后得到的数字出现在[:n-1]最后的位置的概率是（4/5 * 1/4） = 1/5
     
     ... ... 依此类推 ，每次出现在最相对最后位置的概率都是一样的
     
