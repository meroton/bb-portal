#!/usr/bin/env python3

import json
import sys
from pathlib import Path

def main():
    script_dir = Path(__file__).resolve().parent
    api_token_file = script_dir.parent / ".rbe_api_token"
    # Read the API token from a local file.
    api_token = api_token_file.read_text().strip()

    json.dump(
        {
            "headers": {
                "Authorization": [api_token],
            },
        },
        sys.stdout,
    )
    return 0

if __name__ == "__main__":
    sys.exit(main())
