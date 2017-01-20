-module(prob001).
-export([get_sum/0]).
%-compile(export_all).


get_sum() ->
	get_multiples_sum(1, 0).


get_multiples_sum(N, Sum) when N >= 1000 -> Sum;
get_multiples_sum(N, Sum) when N rem 3 == 0; N rem 5 == 0 ->
	get_multiples_sum(N+1, Sum+N);
get_multiples_sum(N, Sum) ->
	get_multiples_sum(N+1, Sum).

