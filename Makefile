rsync:
	rsync -r /media/sf_gohome/src/github.com/awltux/hmrc_oauth ~/go/src/github.com/awltux

git-config:
	git config --global user.name "awltux"; \
	git config --global user.email "chris.d.holman@googlemail.com"; \
	git config --global color.ui true; \
	git config credential.helper store;

# Generate service API code
generate: rsync
	goa gen github.com/awltux/hmrc_oath/design
