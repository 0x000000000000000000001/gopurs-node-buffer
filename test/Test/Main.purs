module Test.Main where

import Prelude
import Effect (Effect)
import Test.Node.Buffer as Buffer

foreign import initFFI :: Unit

main :: Effect Unit
main = Buffer.test
