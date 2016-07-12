package pgengine

import "time"

type Animation struct{
  mFrames []Int
  mIndex Int
  mSPF time.Duration
  mTime Int
  mLoop bool
}

const DefaultSPF = 0.12
func NewAnimation(frames []Int, loop bool, spf time.Duration) *Animation {
  if spf == 0 {
    spf = 0.12 * time.Second
  }
  anim := Animation{
    mFrames: frames,
    mIndex: 1,
    mSPF: spf,
    mTime: 0,
    mLoop: loop
  }

  return &anim
}

func (anim *Animation) Update(dt time.Duration) {
  anim.mTime = m.mTime + df

  if anim.mTime >= m.mSPF {
    anim.mIndex = anim.mIndex + 1
    anim.mTime = 0

    if anim.mIndex > len(anim.mFrames) {
      if anim.mLoop {
        anim.mIndex = 0
      } else {
        anim.mIndex = len(anim.mFrames) - 1
      }
    }
  }
}

func (anim *Animation) SetFrames(frames []Int) {
  anim.mFrames = frames
  anim.mIndex = Min(anim.mIndex, len(anim.mFrames))
}

func (anim *Animation) Frame() Int {
  return anim.mFrames[anim.mIndex]
}

func (anim *Animation) IsFinished() bool {
  return anim.mLoop == false && anim.mIndex == len(anim.mFrames) - 1
}
