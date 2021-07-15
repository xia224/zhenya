# 看到的就记录下来已背后续学习使用

## std::function
Class template std::function is a general-purpose polymorphic function wrapper. Instances of std::function can store, copy, and invoke any CopyConstructible Callable target -- functions, lambda expressions, bind expressions, or other function objects, as well as pointers to member functions and pointers to data members.
白话文：函数变量存储可call的对象，满足拷贝构造和拷贝赋值。

## std::bind
The function template bind generates a forwarding call wrapper for f. Calling this wrapper is equivalent to invoking f with some of its arguments bound to args.
template< class R, class F, class... Args >
/*unspecified*/ bind( F&& f, Args&&... args );
Parameters
f	-	Callable object (function object, pointer to function, reference to function, pointer to member function, or pointer to data member) that will be bound to some arguments
args	-	list of arguments to bind, with the unbound arguments replaced by the placeholders _1, _2, _3... of namespace std::placeholders
白话文：前向函数调用封装，可bound或unbound

## 左值 右值
C++对于左值和右值没有标准定义，但是有一个被广泛认同的说法：
可以取地址的，有名字的，非临时的就是左值；
不能取地址的，没有名字的，临时的就是右值；

C++11 新增右值引用的特性：
类型 && 引用名 = 右值表达式;
在汇编层面右值引用做的事情和常引用是相同的，即产生临时量来存储常量。但是，唯一 一点的区别是，右值引用可以进行读写操作，而常引用只能进行读操作。

## 左右值引用
知乎一篇文章讲的通俗易懂：
[右值引用](https://zhuanlan.zhihu.com/p/97128024)

## 深浅拷贝
todo

## Ranges library (c++20)
std::ranges::dangling  //占位符，标识指针used-after-free
std::ranges::max_element()
std::span()
std::ranges::view

## 编译时类型推到
编译时类型推导，除了我们说过的auto关键字，还有本文的decltype
std::is_same()
static_assert() //编译时静态检查

## c++20新
c++20除了语言的三大特性(concept/coroutine/module)及若干小特性以外，还新增了一些大大小小的库，其中ranges是最大的一个库
