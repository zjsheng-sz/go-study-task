/*
题目2：事务语句
假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
要求 ：
编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
*/

BEGIN TRANSACTION;

DECLARE @balanceA DECIMAL(18,2);

select  @balanceA = balance from accounts where id = 'A'

if @balanceA < 100 
BEGIN
    ROLLBACK TRANSACTION
    PRINT 'A 账户余额不足'
    RETURN;
END


update accounts set balance = (balance - 100) where id = 'A'

update accounts set balance = (balance + 100) where id = 'B'

insert into transactions(from_account_id,to_account_id,amount) VALUES ('A', 'B', 100)

-- 所有操作成功完成，提交事务
COMMIT TRANSACTION