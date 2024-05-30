namespace LRUCache
{
    class LCache<K, V>
    {
        LinkedList<Node<K, V>> list;
        Dictionary<K, LinkedListNode<Node<K, V>>> lookUp;
        readonly int capacity;

        public LCache(int size)
        {
            this.capacity = size;
            list = new LinkedList<Node<K, V>>();
            lookUp = new Dictionary<K, LinkedListNode<Node<K, V>>>(this.capacity);

        }

        public V Get(K key)  {
             
            if(lookUp.TryGetValue(key, out LinkedListNode<Node<K, V>> node))
            {
                list.Remove(node);
                list.AddFirst(node);
                return node.Value.value;
            }

            return default(V);
        }

        public void Put(K k, V v)
        {
            if(lookUp.TryGetValue(k, out LinkedListNode<Node<K, V>> n))
            {
                list.Remove(n);
                n.Value.value = v;
                list.AddFirst(n);
                return;
            }

            if(list.Count >= this.capacity)
            {
                var lastNode = list.Last;
                list.RemoveLast();
                list.AddFirst(new Node<K, V> { key = k, value = v });
                lookUp.Remove(lastNode.Value.key);
            } else
            {
                var newNode = new LinkedListNode<Node<K,V>>(new Node<K,V> { key = k, value = v });
                list.AddFirst(newNode);
                lookUp[k] = newNode;
            }
        }

        public V Front()
        {
            return list.First.Value.value;
        }

        public V Tail()
        {
            return list.Last.Value.value;
        }

    }
}
