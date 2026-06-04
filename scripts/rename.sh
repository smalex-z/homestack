#!/usr/bin/env bash
# One-time setup: rename this template to your own app, then delete this script.
#
# Usage:
#   ./scripts/rename.sh <app-name> [github-owner] ["Display Name"]
#
#   app-name      lowercase identifier — module path, binary, systemd unit (e.g. my-app)
#   github-owner  your GitHub account/org (default: parsed from `git remote origin`)
#   Display Name  shown in the UI and logs (default: derived from app-name)
set -euo pipefail

ROOT="$(cd "$(dirname "$0")/.." && pwd)"
cd "$ROOT"

APP="${1:-}"
if [ -z "$APP" ]; then
    echo "Usage: ./scripts/rename.sh <app-name> [github-owner] [\"Display Name\"]" >&2
    exit 1
fi

# Refuse to run on anything but the pristine template.
if ! grep -q '^module homestack$' go.mod; then
    echo "go.mod module is not 'homestack' — already renamed or not the template root. Nothing to do." >&2
    exit 1
fi

# app-name must be a valid Go module path segment / systemd unit name.
if ! printf '%s' "$APP" | grep -qE '^[a-z][a-z0-9-]*$'; then
    echo "Error: app-name must be lowercase letters, digits, and hyphens, starting with a letter (e.g. my-app)." >&2
    exit 1
fi

# GitHub owner: explicit arg, else parse from the repo's origin remote, else placeholder.
OWNER="${2:-}"
if [ -z "$OWNER" ]; then
    OWNER="$(git config --get remote.origin.url 2>/dev/null \
        | sed -E 's#.*[:/]([^/]+)/[^/]+(\.git)?$#\1#')"
fi
[ -z "$OWNER" ] && OWNER="your-username"

# Display name: explicit arg, else title-case the app-name (my-app -> My App).
TITLE="${3:-}"
if [ -z "$TITLE" ]; then
    TITLE="$(printf '%s' "$APP" | tr '-' ' ' \
        | awk '{for (i=1; i<=NF; i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')"
fi

echo "Renaming template:"
echo "  module / binary : homestack -> $APP"
echo "  display name    : Homestack -> $TITLE"
echo "  github owner    : smalex-z  -> $OWNER"
echo

# Rewrite every file that references the template name. The `smalex-z/homestack`
# substitution runs first so the GitHub owner is fixed before the generic pass.
# This script excludes itself — it's deleted below instead of edited in place.
find . -not -path './.git/*' -not -path './node_modules/*' -not -name 'rename.sh' -type f \
    \( -name '*.go' -o -name '*.sh' -o -name '*.yml' -o -name '*.json' -o -name '*.mod' \
       -o -name '*.ts' -o -name '*.tsx' -o -name '*.html' -o -name 'Makefile' -o -name '.gitignore' \) \
    -exec sed -i \
        -e "s|smalex-z/homestack|$OWNER/$APP|g" \
        -e "s/homestack/$APP/g" \
        -e "s/Homestack/$TITLE/g" {} +

echo "Done. Still to update by hand:"
echo "  - LICENSE          copyright holder and year"
echo "  - README.md        replace with your project's documentation"
echo "  - USING_TEMPLATE.md delete or rewrite for your project"
echo "  - internal/        replace the example users model, service, and routes"
echo "  - frontend/src/    replace the example UI"
echo

rm -- "$0"
echo "Removed scripts/rename.sh."
