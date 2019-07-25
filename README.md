This's an inaccurate (especially not thread-safe) fast PRNG (which is *not* appropriate for cryptographic tasks). It's inaccurate intentionally to make it even faster.

If you need am accurate (but still non-cryptographic) fast PRNG then look at [https://gitlab.com/NebulousLabs/fastrand](https://gitlab.com/NebulousLabs/fastrand).