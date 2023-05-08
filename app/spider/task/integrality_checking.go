package task

func CheckIntegrity() {
	/*
		分为更新与首次录入两种逻辑
		录入机制：
		TODO: 待录入查询机制：1.

		TODO: 1.enterprise_industry by: uscId查询接口 order by updatedTime desc
		TODO: 2.enterprise_info by: uscId查询接口 order by updatedTime desc
		TODO: 3.enterprise_product by: uscId查询接口 order by updatedTime desc
		TODO: 4.enterprise_ranking by: uscId查询接口 order by updatedTime desc

		TODO: 定时任务：回写trades_detail表中公司的数据:通过trades_detail.enterpriseName查询wait_list表中对应的uscId, 通过uscId查询四个数据表.若都存在更改trades_detail表statusCode = 3

		TODO: 定时任务：查询trades_detail表中statusCode=3的行，查询对应数据表回写至trades_detail表中,更新statusCode=4

		TODO: 定时任务：通过contentId查询trades_detail表,若所有记录都为4则更新content的statusCode=3
	*/
}
