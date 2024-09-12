
REPO_SSH="git@github.com:kyleneideck/BackgroundMusic.git"

git clone --no-checkout "$REPO_SSH" ./local
cd local

# 启用 sparse-checkout 模式
git sparse-checkout init --cone
#git sparse-checkout set Images  # 只下载指定的文件夹
git sparse-checkout set BGMDriver/BGMDriver.xcodepro  # 只下载指定的文件夹

# 检出文件
git checkout

# 删除无用文件
find . -maxdepth 1 \
! -name 'BGMDriver/BGMDriver.xcodepro' \
! -name '.git' ! \
-name '.gitignore' \
! -name '.' \
-exec ls -l {} +

#find . -maxdepth 1 ! -name 'Images' ! -name '.git' ! -name '.gitignore' ! -name '.' -exec rm -rf {} +

